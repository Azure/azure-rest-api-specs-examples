from azure.identity import DefaultAzureCredential
from azure.mgmt.mixedreality import MixedRealityClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-mixedreality
# USAGE
    python list_object_anchors_accounts_by_resource_group.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MixedRealityClient(
        credential=DefaultAzureCredential(),
        subscription_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
    )

    response = client.object_anchors_accounts.list_by_resource_group(
        resource_group_name="MyResourceGroup",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/object-anchors/GetByResourceGroup.json
if __name__ == "__main__":
    main()
