from azure.identity import DefaultAzureCredential
from azure.mgmt.frontdoor import FrontDoorManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-frontdoor
# USAGE
    python frontdoor_rules_engine_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = FrontDoorManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.rules_engines.begin_create_or_update(
        resource_group_name="rg1",
        front_door_name="frontDoor1",
        rules_engine_name="rulesEngine1",
        rules_engine_parameters={
            "properties": {
                "rules": [
                    {
                        "action": {
                            "routeConfigurationOverride": {
                                "@odata.type": "#Microsoft.Azure.FrontDoor.Models.FrontdoorRedirectConfiguration",
                                "customFragment": "fragment",
                                "customHost": "www.bing.com",
                                "customPath": "/api",
                                "customQueryString": "a=b",
                                "redirectProtocol": "HttpsOnly",
                                "redirectType": "Moved",
                            }
                        },
                        "matchConditions": [
                            {
                                "rulesEngineMatchValue": ["CH"],
                                "rulesEngineMatchVariable": "RemoteAddr",
                                "rulesEngineOperator": "GeoMatch",
                            }
                        ],
                        "matchProcessingBehavior": "Stop",
                        "name": "Rule1",
                        "priority": 1,
                    },
                    {
                        "action": {
                            "responseHeaderActions": [
                                {
                                    "headerActionType": "Overwrite",
                                    "headerName": "Cache-Control",
                                    "value": "public, max-age=31536000",
                                }
                            ]
                        },
                        "matchConditions": [
                            {
                                "rulesEngineMatchValue": ["jpg"],
                                "rulesEngineMatchVariable": "RequestFilenameExtension",
                                "rulesEngineOperator": "Equal",
                                "transforms": ["Lowercase"],
                            }
                        ],
                        "name": "Rule2",
                        "priority": 2,
                    },
                    {
                        "action": {
                            "routeConfigurationOverride": {
                                "@odata.type": "#Microsoft.Azure.FrontDoor.Models.FrontdoorForwardingConfiguration",
                                "backendPool": {
                                    "id": "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/backendPools/backendPool1"
                                },
                                "cacheConfiguration": {
                                    "cacheDuration": "P1DT12H20M30S",
                                    "dynamicCompression": "Disabled",
                                    "queryParameterStripDirective": "StripOnly",
                                    "queryParameters": "a=b,p=q",
                                },
                                "customForwardingPath": None,
                                "forwardingProtocol": "HttpsOnly",
                            }
                        },
                        "matchConditions": [
                            {
                                "negateCondition": False,
                                "rulesEngineMatchValue": ["allowoverride"],
                                "rulesEngineMatchVariable": "RequestHeader",
                                "rulesEngineOperator": "Equal",
                                "selector": "Rules-Engine-Route-Forward",
                                "transforms": ["Lowercase"],
                            }
                        ],
                        "name": "Rule3",
                        "priority": 3,
                    },
                ]
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/FrontdoorRulesEngineCreate.json
if __name__ == "__main__":
    main()
