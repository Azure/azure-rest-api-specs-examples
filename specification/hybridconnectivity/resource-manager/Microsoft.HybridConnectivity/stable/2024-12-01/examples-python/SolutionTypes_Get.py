from azure.identity import DefaultAzureCredential

from azure.mgmt.hybridconnectivity import HybridConnectivityMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hybridconnectivity
# USAGE
    python solution_types_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HybridConnectivityMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.solution_types.get(
        resource_group_name="rgpublicCloud",
        solution_type="lulzqllpu",
    )
    print(response)


# x-ms-original-file: 2024-12-01/SolutionTypes_Get.json
if __name__ == "__main__":
    main()
