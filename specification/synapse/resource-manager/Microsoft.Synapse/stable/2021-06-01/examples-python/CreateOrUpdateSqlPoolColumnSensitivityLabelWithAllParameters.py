from azure.identity import DefaultAzureCredential
from azure.mgmt.synapse import SynapseManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-synapse
# USAGE
    python create_or_update_sql_pool_column_sensitivity_label_with_all_parameters.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SynapseManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-1111-2222-3333-444444444444",
    )

    response = client.sql_pool_sensitivity_labels.create_or_update(
        resource_group_name="myRG",
        workspace_name="myServer",
        sql_pool_name="myDatabase",
        schema_name="dbo",
        table_name="myTable",
        column_name="myColumn",
        parameters={
            "properties": {
                "informationType": "PhoneNumber",
                "informationTypeId": "d22fa6e9-5ee4-3bde-4c2b-a409604c4646",
                "labelId": "bf91e08c-f4f0-478a-b016-25164b2a65ff",
                "labelName": "PII",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateSqlPoolColumnSensitivityLabelWithAllParameters.json
if __name__ == "__main__":
    main()
