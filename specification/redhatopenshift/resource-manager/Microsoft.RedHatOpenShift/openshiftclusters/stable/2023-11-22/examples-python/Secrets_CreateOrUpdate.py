from azure.identity import DefaultAzureCredential

from azure.mgmt.redhatopenshift import AzureRedHatOpenShiftClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-redhatopenshift
# USAGE
    python secrets_create_or_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureRedHatOpenShiftClient(
        credential=DefaultAzureCredential(),
        subscription_id="subscriptionId",
    )

    response = client.secrets.create_or_update(
        resource_group_name="resourceGroup",
        resource_name="resourceName",
        child_resource_name="childResourceName",
        parameters={"properties": {}},
    )
    print(response)


# x-ms-original-file: specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/openshiftclusters/stable/2023-11-22/examples/Secrets_CreateOrUpdate.json
if __name__ == "__main__":
    main()
