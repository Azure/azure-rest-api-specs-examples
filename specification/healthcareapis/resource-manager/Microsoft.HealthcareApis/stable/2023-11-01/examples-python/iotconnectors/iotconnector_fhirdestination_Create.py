from azure.identity import DefaultAzureCredential
from azure.mgmt.healthcareapis import HealthcareApisManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-healthcareapis
# USAGE
    python iotconnector_fhirdestination_create.py

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

    response = client.iot_connector_fhir_destination.begin_create_or_update(
        resource_group_name="testRG",
        workspace_name="workspace1",
        iot_connector_name="blue",
        fhir_destination_name="dest1",
        iot_fhir_destination={
            "location": "westus",
            "properties": {
                "fhirMapping": {
                    "content": {
                        "template": [
                            {
                                "template": {
                                    "codes": [
                                        {"code": "8867-4", "display": "Heart rate", "system": "http://loinc.org"}
                                    ],
                                    "periodInterval": 60,
                                    "typeName": "heartrate",
                                    "value": {
                                        "defaultPeriod": 5000,
                                        "unit": "count/min",
                                        "valueName": "hr",
                                        "valueType": "SampledData",
                                    },
                                },
                                "templateType": "CodeValueFhir",
                            }
                        ],
                        "templateType": "CollectionFhirTemplate",
                    }
                },
                "fhirServiceResourceId": "subscriptions/11111111-2222-3333-4444-555566667777/resourceGroups/myrg/providers/Microsoft.HealthcareApis/workspaces/myworkspace/fhirservices/myfhirservice",
                "resourceIdentityResolutionType": "Create",
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2023-11-01/examples/iotconnectors/iotconnector_fhirdestination_Create.json
if __name__ == "__main__":
    main()
