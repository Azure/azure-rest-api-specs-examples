from azure.identity import DefaultAzureCredential

from azure.mgmt.apimanagement import ApiManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-apimanagement
# USAGE
    python api_management_create_workspace_backend_proxy_backend.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ApiManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.workspace_backend.create_or_update(
        resource_group_name="rg1",
        service_name="apimService1",
        workspace_id="wks1",
        backend_id="proxybackend",
        parameters={
            "properties": {
                "credentials": {
                    "authorization": {"parameter": "opensesma", "scheme": "Basic"},
                    "header": {"x-my-1": ["val1", "val2"]},
                    "query": {"sv": ["xx", "bb", "cc"]},
                },
                "description": "description5308",
                "protocol": "http",
                "proxy": {"password": "<password>", "url": "http://192.168.1.1:8080", "username": "Contoso\\admin"},
                "tls": {"validateCertificateChain": True, "validateCertificateName": True},
                "url": "https://backendname2644/",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateWorkspaceBackendProxyBackend.json
if __name__ == "__main__":
    main()
