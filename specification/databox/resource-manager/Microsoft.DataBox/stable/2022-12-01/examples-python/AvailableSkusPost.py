from azure.identity import DefaultAzureCredential
from azure.mgmt.databox import DataBoxManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-databox
# USAGE
    python available_skus_post.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DataBoxManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="YourSubscriptionId",
    )

    response = client.service.list_available_skus_by_resource_group(
        resource_group_name="YourResourceGroupName",
        location="westus",
        available_sku_request={"country": "XX", "location": "westus", "transferType": "ImportToAzure"},
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/AvailableSkusPost.json
if __name__ == "__main__":
    main()
