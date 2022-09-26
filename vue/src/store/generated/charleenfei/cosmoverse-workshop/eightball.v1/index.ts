import { txClient, queryClient, MissingWalletError , registry} from './module'

import { Fortune } from "./module/types/eightball/v1/fortunes"
import { FortuneList } from "./module/types/eightball/v1/fortunes"
import { Params } from "./module/types/eightball/v1/params"
import { QueryParamsRequest } from "./module/types/eightball/v1/query"
import { QueryParamsResponse } from "./module/types/eightball/v1/query"
import { MsgSwap } from "./module/types/eightball/v1/swap"
import { MsgSwapResponse } from "./module/types/eightball/v1/swap"
import { Workflow } from "./module/types/eightball/v1/workflow"


export { Fortune, FortuneList, Params, QueryParamsRequest, QueryParamsResponse, MsgSwap, MsgSwapResponse, Workflow };

async function initTxClient(vuexGetters) {
	return await txClient(vuexGetters['common/wallet/signer'], {
		addr: vuexGetters['common/env/apiTendermint']
	})
}

async function initQueryClient(vuexGetters) {
	return await queryClient({
		addr: vuexGetters['common/env/apiCosmos']
	})
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

function getStructure(template) {
	let structure = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field: any = {}
		field.name = key
		field.type = typeof value
		structure.fields.push(field)
	}
	return structure
}

const getDefaultState = () => {
	return {
				Fortune: {},
				Fortunes: {},
				
				_Structure: {
						Fortune: getStructure(Fortune.fromPartial({})),
						FortuneList: getStructure(FortuneList.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						QueryParamsRequest: getStructure(QueryParamsRequest.fromPartial({})),
						QueryParamsResponse: getStructure(QueryParamsResponse.fromPartial({})),
						MsgSwap: getStructure(MsgSwap.fromPartial({})),
						MsgSwapResponse: getStructure(MsgSwapResponse.fromPartial({})),
						Workflow: getStructure(Workflow.fromPartial({})),
						
		},
		_Registry: registry,
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(JSON.stringify(subscription))
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(JSON.stringify(subscription))
		}
	},
	getters: {
				getFortune: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Fortune[JSON.stringify(params)] ?? {}
		},
				getFortunes: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Fortunes[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: eightball.v1 initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					const sub=JSON.parse(subscription)
					await dispatch(sub.action, sub.payload)
				}catch(e) {
					throw new Error('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryFortune({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryFortune( key.owner)).data
				
					
				commit('QUERY', { query: 'Fortune', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryFortune', payload: { options: { all }, params: {...key},query }})
				return getters['getFortune']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryFortune API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryFortunes({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryFortunes()).data
				
					
				commit('QUERY', { query: 'Fortunes', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryFortunes', payload: { options: { all }, params: {...key},query }})
				return getters['getFortunes']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryFortunes API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgConnectToDex({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgConnectToDex(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgConnectToDex:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgConnectToDex:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgFeelingLucky({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgFeelingLucky(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgFeelingLucky:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgFeelingLucky:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgConnectToDex({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgConnectToDex(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgConnectToDex:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgConnectToDex:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgFeelingLucky({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgFeelingLucky(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgFeelingLucky:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgFeelingLucky:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
