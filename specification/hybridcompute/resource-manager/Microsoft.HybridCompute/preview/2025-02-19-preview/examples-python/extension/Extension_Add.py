from azure.identity import DefaultAzureCredential

from azure.mgmt.hybridcompute import HybridComputeManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hybridcompute
# USAGE
    python extension_add.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HybridComputeManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="{subscriptionId}",
    )

    response = client.begin_setup_extensions(
        resource_group_name="myResourceGroup",
        machine_name="myMachine",
        extensions={
            "extensions": [
                {"publisher": "Microsoft.Azure.Monitoring", "type": "AzureMonitorAgentLinux"},
                {"publisher": "<extension_publisher>", "type": "<extension_type>"},
            ]
        },
    ).result()
    print(response)


# x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/extension/Extension_Add.json
if __name__ == "__main__":
    main()
