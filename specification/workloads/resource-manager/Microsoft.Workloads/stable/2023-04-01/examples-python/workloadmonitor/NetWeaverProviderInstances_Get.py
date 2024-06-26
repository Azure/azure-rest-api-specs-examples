from azure.identity import DefaultAzureCredential
from azure.mgmt.workloads import WorkloadsMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-workloads
# USAGE
    python net_weaver_provider_instances_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = WorkloadsMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.provider_instances.get(
        resource_group_name="myResourceGroup",
        monitor_name="mySapMonitor",
        provider_instance_name="myProviderInstance",
    )
    print(response)


# x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/NetWeaverProviderInstances_Get.json
if __name__ == "__main__":
    main()
