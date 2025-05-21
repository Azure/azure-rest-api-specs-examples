from azure.identity import DefaultAzureCredential

from azure.mgmt.resource.databoundaries import DataBoundaryMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-resource
# USAGE
    python get_tenant_data_boundary.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DataBoundaryMgmtClient(
        credential=DefaultAzureCredential(),
    )

    response = client.data_boundaries.get_tenant(
        default="default",
    )
    print(response)


# x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2024-08-01/examples/GetTenantDataBoundary.json
if __name__ == "__main__":
    main()
