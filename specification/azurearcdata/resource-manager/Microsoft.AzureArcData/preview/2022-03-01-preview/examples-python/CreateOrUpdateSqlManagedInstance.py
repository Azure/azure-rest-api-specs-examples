from azure.identity import DefaultAzureCredential
from azure.mgmt.azurearcdata import AzureArcDataManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-azurearcdata
# USAGE
    python create_or_update_sql_managed_instance.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureArcDataManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-1111-2222-3333-444444444444",
    )

    response = client.sql_managed_instances.begin_create(
        resource_group_name="testrg",
        sql_managed_instance_name="testsqlManagedInstance",
        sql_managed_instance={
            "extendedLocation": {
                "name": "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/arclocation",
                "type": "CustomLocation",
            },
            "location": "northeurope",
            "properties": {
                "activeDirectoryInformation": {"keytabInformation": {"keytab": "********"}},
                "admin": "Admin user",
                "basicLoginInformation": {"password": "********", "username": "username"},
                "clusterId": "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s",
                "endTime": "Instance end time",
                "extensionId": "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s/providers/Microsoft.KubernetesConfiguration/extensions/extension",
                "k8sRaw": {
                    "additionalProperty": 1234,
                    "spec": {
                        "replicas": 1,
                        "scheduling": {
                            "default": {
                                "resources": {
                                    "limits": {"additionalProperty": "additionalValue", "cpu": "1", "memory": "8Gi"},
                                    "requests": {"additionalProperty": "additionalValue", "cpu": "1", "memory": "8Gi"},
                                }
                            }
                        },
                    },
                },
                "licenseType": "LicenseIncluded",
                "startTime": "Instance start time",
            },
            "sku": {"dev": True, "name": "vCore", "tier": "GeneralPurpose"},
            "tags": {"mytag": "myval"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/preview/2022-03-01-preview/examples/CreateOrUpdateSqlManagedInstance.json
if __name__ == "__main__":
    main()
