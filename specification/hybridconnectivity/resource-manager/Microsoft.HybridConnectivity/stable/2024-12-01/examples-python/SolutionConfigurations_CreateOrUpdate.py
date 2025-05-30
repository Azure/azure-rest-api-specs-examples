from azure.identity import DefaultAzureCredential

from azure.mgmt.hybridconnectivity import HybridConnectivityMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hybridconnectivity
# USAGE
    python solution_configurations_create_or_update.py

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

    response = client.solution_configurations.create_or_update(
        resource_uri="ymuj",
        solution_configuration="keebwujt",
        resource={"properties": {"solutionSettings": {}, "solutionType": "nmtqllkyohwtsthxaimsye"}},
    )
    print(response)


# x-ms-original-file: 2024-12-01/SolutionConfigurations_CreateOrUpdate.json
if __name__ == "__main__":
    main()
