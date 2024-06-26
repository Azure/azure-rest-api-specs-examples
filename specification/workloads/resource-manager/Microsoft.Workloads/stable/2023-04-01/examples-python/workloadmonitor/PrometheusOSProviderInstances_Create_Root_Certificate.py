from azure.identity import DefaultAzureCredential
from azure.mgmt.workloads import WorkloadsMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-workloads
# USAGE
    python prometheus_os_provider_instances_create_root_certificate.py

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

    response = client.provider_instances.begin_create(
        resource_group_name="myResourceGroup",
        monitor_name="mySapMonitor",
        provider_instance_name="myProviderInstance",
        provider_instance_parameter={
            "properties": {
                "providerSettings": {
                    "prometheusUrl": "http://192.168.0.0:9090/metrics",
                    "providerType": "PrometheusOS",
                    "sapSid": "SID",
                    "sslPreference": "RootCertificate",
                }
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/PrometheusOSProviderInstances_Create_Root_Certificate.json
if __name__ == "__main__":
    main()
