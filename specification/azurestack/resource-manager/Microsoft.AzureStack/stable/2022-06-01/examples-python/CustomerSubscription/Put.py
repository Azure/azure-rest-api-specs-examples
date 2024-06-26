from azure.identity import DefaultAzureCredential
from azure.mgmt.azurestack import AzureStackManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-azurestack
# USAGE
    python put.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureStackManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="dd8597b4-8739-4467-8b10-f8679f62bfbf",
    )

    response = client.customer_subscriptions.create(
        resource_group="azurestack",
        registration_name="testregistration",
        customer_subscription_name="E09A4E93-29A7-4EBA-A6D4-76202383F07F",
        customer_creation_parameters={"properties": {"tenantId": "dbab3982-796f-4d03-9908-044c08aef8a2"}},
    )
    print(response)


# x-ms-original-file: specification/azurestack/resource-manager/Microsoft.AzureStack/stable/2022-06-01/examples/CustomerSubscription/Put.json
if __name__ == "__main__":
    main()
