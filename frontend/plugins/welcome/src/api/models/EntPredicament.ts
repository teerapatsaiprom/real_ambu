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
    EntPredicamentEdges,
    EntPredicamentEdgesFromJSON,
    EntPredicamentEdgesFromJSONTyped,
    EntPredicamentEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPredicament
 */
export interface EntPredicament {
    /**
     * AddedTime holds the value of the "Added_Time" field.
     * @type {string}
     * @memberof EntPredicament
     */
    addedTime?: string;
    /**
     * 
     * @type {EntPredicamentEdges}
     * @memberof EntPredicament
     */
    edges?: EntPredicamentEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntPredicament
     */
    id?: number;
}

export function EntPredicamentFromJSON(json: any): EntPredicament {
    return EntPredicamentFromJSONTyped(json, false);
}

export function EntPredicamentFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPredicament {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'addedTime': !exists(json, 'Added_Time') ? undefined : json['Added_Time'],
        'edges': !exists(json, 'edges') ? undefined : EntPredicamentEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntPredicamentToJSON(value?: EntPredicament | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Added_Time': value.addedTime,
        'edges': EntPredicamentEdgesToJSON(value.edges),
        'id': value.id,
    };
}


