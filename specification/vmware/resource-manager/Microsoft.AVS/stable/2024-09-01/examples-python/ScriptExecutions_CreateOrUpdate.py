from azure.identity import DefaultAzureCredential

from azure.mgmt.avs import AVSClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-avs
# USAGE
    python script_executions_create_or_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AVSClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.script_executions.begin_create_or_update(
        resource_group_name="group1",
        private_cloud_name="cloud1",
        script_execution_name="addSsoServer",
        script_execution={
            "properties": {
                "hiddenParameters": [{"name": "Password", "secureValue": "PlaceholderPassword", "type": "SecureValue"}],
                "parameters": [
                    {"name": "DomainName", "type": "Value", "value": "placeholderDomain.local"},
                    {"name": "BaseUserDN", "type": "Value", "value": "DC=placeholder, DC=placeholder"},
                ],
                "retention": "P0Y0M60DT0H60M60S",
                "scriptCmdletId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/scriptPackages/AVS.PowerCommands@1.0.0/scriptCmdlets/New-SsoExternalIdentitySource",
                "timeout": "P0Y0M0DT0H60M60S",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2024-09-01/examples/ScriptExecutions_CreateOrUpdate.json
if __name__ == "__main__":
    main()
