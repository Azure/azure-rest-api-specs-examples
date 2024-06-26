from azure.identity import DefaultAzureCredential
from azure.mgmt.workloads import WorkloadsMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-workloads
# USAGE
    python sap_application_server_instances_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = WorkloadsMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="6d875e77-e412-4d7d-9af4-8895278b4443",
    )

    response = client.sap_application_server_instances.get(
        resource_group_name="test-rg",
        sap_virtual_instance_name="X00",
        application_instance_name="app01",
    )
    print(response)


# x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPApplicationServerInstances_Get.json
if __name__ == "__main__":
    main()
