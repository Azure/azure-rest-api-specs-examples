from azure.identity import DefaultAzureCredential
from azure.mgmt.agrifood import AgriFoodMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-agrifood
# USAGE
    python farm_beats_models_get_operation_result.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AgriFoodMgmtClient(
        credential=DefaultAzureCredential(),
        solution_id="SOLUTION_ID",
        subscription_id="11111111-2222-3333-4444-555555555555",
    )

    response = client.farm_beats_models.get_operation_result(
        resource_group_name="examples-rg",
        farm_beats_resource_name="examples-farmBeatsResourceName",
        operation_results_id="resource-provisioning-id-farmBeatsResourceName",
    )
    print(response)


# x-ms-original-file: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/FarmBeatsModels_GetOperationResult.json
if __name__ == "__main__":
    main()
