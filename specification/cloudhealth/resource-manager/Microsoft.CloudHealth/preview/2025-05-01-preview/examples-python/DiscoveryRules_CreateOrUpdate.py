from azure.identity import DefaultAzureCredential

from azure.mgmt.cloudhealth import CloudHealthMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-cloudhealth
# USAGE
    python discovery_rules_create_or_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = CloudHealthMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.discovery_rules.create_or_update(
        resource_group_name="myResourceGroup",
        health_model_name="myHealthModel",
        discovery_rule_name="myDiscoveryRule",
        resource={
            "properties": {
                "addRecommendedSignals": "Enabled",
                "authenticationSetting": "authSetting1",
                "discoverRelationships": "Enabled",
                "displayName": "myDisplayName",
                "resourceGraphQuery": "resources | where subscriptionId == '7ddfffd7-9b32-40df-1234-828cbd55d6f4' | where resourceGroup == 'my-rg'",
            }
        },
    )
    print(response)


# x-ms-original-file: 2025-05-01-preview/DiscoveryRules_CreateOrUpdate.json
if __name__ == "__main__":
    main()
