from azure.identity import DefaultAzureCredential

from azure.mgmt.lambdatesthyperexecute import LambdaTestHyperExecuteMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-lambdatesthyperexecute
# USAGE
    python organizations_list_by_resource_group_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = LambdaTestHyperExecuteMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.organizations.list_by_resource_group(
        resource_group_name="rgopenapi",
    )
    for item in response:
        print(item)


# x-ms-original-file: 2024-02-01/Organizations_ListByResourceGroup_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
