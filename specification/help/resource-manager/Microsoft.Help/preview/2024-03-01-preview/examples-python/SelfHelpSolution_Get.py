from azure.identity import DefaultAzureCredential

from azure.mgmt.selfhelp import SelfHelpMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-selfhelp
# USAGE
    python self_help_solution_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SelfHelpMgmtClient(
        credential=DefaultAzureCredential(),
    )

    response = client.solution_self_help.get(
        solution_id="SolutionId1",
    )
    print(response)


# x-ms-original-file: specification/help/resource-manager/Microsoft.Help/preview/2024-03-01-preview/examples/SelfHelpSolution_Get.json
if __name__ == "__main__":
    main()
