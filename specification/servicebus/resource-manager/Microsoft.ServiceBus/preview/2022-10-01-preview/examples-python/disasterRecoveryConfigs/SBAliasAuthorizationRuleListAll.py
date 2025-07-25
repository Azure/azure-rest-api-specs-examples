from azure.identity import DefaultAzureCredential

from azure.mgmt.servicebus import ServiceBusManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-servicebus
# USAGE
    python sb_alias_authorization_rule_list_all.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ServiceBusManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="exampleSubscriptionId",
    )

    response = client.disaster_recovery_configs.list_authorization_rules(
        resource_group_name="exampleResourceGroup",
        namespace_name="sdk-Namespace-9080",
        alias="sdk-DisasterRecovery-4047",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/disasterRecoveryConfigs/SBAliasAuthorizationRuleListAll.json
if __name__ == "__main__":
    main()
