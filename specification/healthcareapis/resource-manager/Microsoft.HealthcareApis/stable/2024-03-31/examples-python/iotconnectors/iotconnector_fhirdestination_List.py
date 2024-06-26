from azure.identity import DefaultAzureCredential

from azure.mgmt.healthcareapis import HealthcareApisManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-healthcareapis
# USAGE
    python iotconnector_fhirdestination_list.py

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

    response = client.fhir_destinations.list_by_iot_connector(
        resource_group_name="testRG",
        workspace_name="workspace1",
        iot_connector_name="blue",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/iotconnectors/iotconnector_fhirdestination_List.json
if __name__ == "__main__":
    main()
