from azure.identity import DefaultAzureCredential

from azure.mgmt.databox import DataBoxManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-databox
# USAGE
    python job_mitigate.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DataBoxManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="YourSubscriptionId",
    )

    client.mitigate(
        job_name="TestJobName1",
        resource_group_name="YourResourceGroupName",
        mitigate_job_request={
            "serialNumberCustomerResolutionMap": {"testDISK-1": "MoveToCleanUpDevice", "testDISK-2": "Resume"}
        },
    )


# x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2025-02-01/examples/JobMitigate.json
if __name__ == "__main__":
    main()
