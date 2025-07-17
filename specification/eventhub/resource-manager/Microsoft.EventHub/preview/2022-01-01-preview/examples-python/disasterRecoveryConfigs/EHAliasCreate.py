from azure.identity import DefaultAzureCredential

from azure.mgmt.eventhub import EventHubManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-eventhub
# USAGE
    python eh_alias_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = EventHubManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="exampleSubscriptionId",
    )

    response = client.disaster_recovery_configs.create_or_update(
        resource_group_name="exampleResourceGroup",
        namespace_name="sdk-Namespace-8859",
        alias="sdk-DisasterRecovery-3814",
        parameters={"properties": {"partnerNamespace": "sdk-Namespace-37"}},
    )
    print(response)


# x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-01-01-preview/examples/disasterRecoveryConfigs/EHAliasCreate.json
if __name__ == "__main__":
    main()
