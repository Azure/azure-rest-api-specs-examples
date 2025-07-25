from azure.identity import DefaultAzureCredential

from azure.mgmt.web import WebSiteManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-web
# USAGE
    python create_or_update_static_site.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = WebSiteManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="34adfa4f-cedf-4dc0-ba29-b6d1a69ab345",
    )

    response = client.static_sites.begin_create_or_update_static_site(
        resource_group_name="rg",
        name="testStaticSite0",
        static_site_envelope={
            "location": "West US 2",
            "properties": {
                "branch": "master",
                "buildProperties": {"apiLocation": "api", "appArtifactLocation": "build", "appLocation": "app"},
                "repositoryToken": "repoToken123",
                "repositoryUrl": "https://github.com/username/RepoName",
            },
            "sku": {"name": "Basic", "tier": "Basic"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/CreateOrUpdateStaticSite.json
if __name__ == "__main__":
    main()
