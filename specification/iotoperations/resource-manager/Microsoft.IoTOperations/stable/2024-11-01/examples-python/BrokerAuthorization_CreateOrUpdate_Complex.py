from azure.identity import DefaultAzureCredential

from azure.mgmt.iotoperations import IoTOperationsMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-iotoperations
# USAGE
    python broker_authorization_create_or_update_complex.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = IoTOperationsMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.broker_authorization.begin_create_or_update(
        resource_group_name="rgiotoperations",
        instance_name="resource-name123",
        broker_name="resource-name123",
        authorization_name="resource-name123",
        resource={
            "extendedLocation": {"name": "qmbrfwcpwwhggszhrdjv", "type": "CustomLocation"},
            "properties": {
                "authorizationPolicies": {
                    "cache": "Enabled",
                    "rules": [
                        {
                            "brokerResources": [
                                {"clientIds": ["{principal.attributes.building}*"], "method": "Connect"},
                                {
                                    "method": "Publish",
                                    "topics": [
                                        "sensors/{principal.attributes.building}/{principal.clientId}/telemetry/*"
                                    ],
                                },
                                {"method": "Subscribe", "topics": ["commands/{principal.attributes.organization}"]},
                            ],
                            "principals": {
                                "attributes": [{"building": "17", "organization": "contoso"}],
                                "usernames": ["temperature-sensor", "humidity-sensor"],
                            },
                            "stateStoreResources": [
                                {
                                    "keyType": "Pattern",
                                    "keys": [
                                        "myreadkey",
                                        "myotherkey?",
                                        "mynumerickeysuffix[0-9]",
                                        "clients:{principal.clientId}:*",
                                    ],
                                    "method": "Read",
                                },
                                {"keyType": "Binary", "keys": ["MTE2IDEwMSAxMTUgMTE2"], "method": "ReadWrite"},
                            ],
                        }
                    ],
                }
            },
        },
    ).result()
    print(response)


# x-ms-original-file: 2024-11-01/BrokerAuthorization_CreateOrUpdate_Complex.json
if __name__ == "__main__":
    main()
