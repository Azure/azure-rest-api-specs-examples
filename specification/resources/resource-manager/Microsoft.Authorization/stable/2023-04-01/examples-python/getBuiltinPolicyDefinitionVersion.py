from azure.identity import DefaultAzureCredential

from azure.mgmt.resource.policy import PolicyClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-resource
# USAGE
    python get_builtin_policy_definition_version.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = PolicyClient(
        credential=DefaultAzureCredential(),
        subscription_id="ae640e6b-ba3e-4256-9d62-2993eecfa6f2",
    )

    response = client.policy_definition_versions.get_built_in(
        policy_definition_name="7433c107-6db4-4ad1-b57a-a76dce0154a1",
        policy_definition_version="1.2.1",
    )
    print(response)


# x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/getBuiltinPolicyDefinitionVersion.json
if __name__ == "__main__":
    main()
