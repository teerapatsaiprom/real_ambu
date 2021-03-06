/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntPredicament,
    EntPredicamentFromJSON,
    EntPredicamentFromJSONTyped,
    EntPredicamentToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntStatuscarEdges
 */
export interface EntStatuscarEdges {
    /**
     * StatuscarPredicament holds the value of the statuscar_predicament edge.
     * @type {Array<EntPredicament>}
     * @memberof EntStatuscarEdges
     */
    statuscarPredicament?: Array<EntPredicament>;
}

export function EntStatuscarEdgesFromJSON(json: any): EntStatuscarEdges {
    return EntStatuscarEdgesFromJSONTyped(json, false);
}

export function EntStatuscarEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntStatuscarEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'statuscarPredicament': !exists(json, 'statuscarPredicament') ? undefined : ((json['statuscarPredicament'] as Array<any>).map(EntPredicamentFromJSON)),
    };
}

export function EntStatuscarEdgesToJSON(value?: EntStatuscarEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'statuscarPredicament': value.statuscarPredicament === undefined ? undefined : ((value.statuscarPredicament as Array<any>).map(EntPredicamentToJSON)),
    };
}


