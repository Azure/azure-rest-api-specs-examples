from azure.identity import DefaultAzureCredential
from azure.mgmt.workloadssapvirtualinstance import WorkloadsSapVirtualInstanceMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-workloadssapvirtualinstance
# USAGE
    python sap_virtual_instances_detect_installation_ha_av_zone.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = WorkloadsSapVirtualInstanceMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="8e17e36c-42e9-4cd5-a078-7b44883414e0",
    )

    response = client.sap_virtual_instances.begin_create(
        resource_group_name="test-rg",
        sap_virtual_instance_name="X00",
    ).result()
    print(response)


# x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapvirtualinstances/SAPVirtualInstances_DetectInstallation_HA_AvZone.json
if __name__ == "__main__":
    main()
