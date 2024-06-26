from azure.identity import DefaultAzureCredential

from azure.mgmt.healthcareapis import HealthcareApisManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-healthcareapis
# USAGE
    python dicom_services_delete.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HealthcareApisManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    client.dicom_services.begin_delete(
        resource_group_name="testRG",
        dicom_service_name="blue",
        workspace_name="workspace1",
    ).result()


# x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/dicomservices/DicomServices_Delete.json
if __name__ == "__main__":
    main()
