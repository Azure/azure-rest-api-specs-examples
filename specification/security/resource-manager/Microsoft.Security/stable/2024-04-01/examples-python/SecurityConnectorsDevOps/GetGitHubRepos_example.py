from azure.identity import DefaultAzureCredential

from azure.mgmt.security import SecurityCenter

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-security
# USAGE
    python get_git_hub_repos_example.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SecurityCenter(
        credential=DefaultAzureCredential(),
        subscription_id="0806e1cd-cfda-4ff8-b99c-2b0af42cffd3",
    )

    response = client.git_hub_repos.get(
        resource_group_name="myRg",
        security_connector_name="mySecurityConnectorName",
        owner_name="myGitHubOwner",
        repo_name="myGitHubRepo",
    )
    print(response)


# x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2024-04-01/examples/SecurityConnectorsDevOps/GetGitHubRepos_example.json
if __name__ == "__main__":
    main()
