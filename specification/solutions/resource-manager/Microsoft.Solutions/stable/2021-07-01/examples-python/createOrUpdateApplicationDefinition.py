from azure.identity import DefaultAzureCredential
from azure.mgmt.managedapplications import ManagedApplicationsMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-managedapplications
# USAGE
    python create_or_update_application_definition.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ManagedApplicationsMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.application_definitions.create_or_update(
        resource_group_name="rg",
        application_definition_name="myManagedApplicationDef",
        parameters={
            "properties": {
                "authorizations": [{"principalId": "validprincipalguid", "roleDefinitionId": "validroleguid"}],
                "description": "myManagedApplicationDef description",
                "displayName": "myManagedApplicationDef",
                "lockLevel": "None",
                "packageFileUri": "https://path/to/packagezipfile",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/createOrUpdateApplicationDefinition.json
if __name__ == "__main__":
    main()
