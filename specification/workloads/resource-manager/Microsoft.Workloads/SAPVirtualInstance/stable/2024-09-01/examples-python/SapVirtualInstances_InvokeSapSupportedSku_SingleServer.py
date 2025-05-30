from azure.identity import DefaultAzureCredential

from azure.mgmt.workloadssapvirtualinstance import WorkloadsSapVirtualInstanceMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-workloadssapvirtualinstance
# USAGE
    python sap_virtual_instances_invoke_sap_supported_sku_single_server.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = WorkloadsSapVirtualInstanceMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.sap_virtual_instances.get_sap_supported_sku(
        location="centralus",
        body={
            "appLocation": "eastus",
            "databaseType": "HANA",
            "deploymentType": "SingleServer",
            "environment": "NonProd",
            "sapProduct": "S4HANA",
        },
    )
    print(response)


# x-ms-original-file: 2024-09-01/SapVirtualInstances_InvokeSapSupportedSku_SingleServer.json
if __name__ == "__main__":
    main()
