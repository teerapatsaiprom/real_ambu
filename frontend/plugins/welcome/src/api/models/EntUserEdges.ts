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
 * @interface EntUserEdges
 */
export interface EntUserEdges {
    /**
     * UserPredicament holds the value of the user_predicament edge.
     * @type {Array<EntPredicament>}
     * @memberof EntUserEdges
     */
    userPredicament?: Array<EntPredicament>;
}

export function EntUserEdgesFromJSON(json: any): EntUserEdges {
    return EntUserEdgesFromJSONTyped(json, false);
}

export function EntUserEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntUserEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'userPredicament': !exists(json, 'userPredicament') ? undefined : ((json['userPredicament'] as Array<any>).map(EntPredicamentFromJSON)),
    };
}

export function EntUserEdgesToJSON(value?: EntUserEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'userPredicament': value.userPredicament === undefined ? undefined : ((value.userPredicament as Array<any>).map(EntPredicamentToJSON)),
    };
}


