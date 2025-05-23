from azure.identity import DefaultAzureCredential

from azure.mgmt.migrationassessment import MigrationAssessmentMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-migrationassessment
# USAGE
    python avs_assessments_operations_create_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MigrationAssessmentMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="4bd2aa0f-2bd2-4d67-91a8-5a4533d58600",
    )

    response = client.avs_assessments_operations.begin_create(
        resource_group_name="ayagrawrg",
        project_name="app18700project",
        group_name="kuchatur-test",
        assessment_name="asm2",
        resource={
            "properties": {
                "assessmentType": "AvsAssessment",
                "azureLocation": "EastUs",
                "azureOfferCode": "MSAZR0003P",
                "currency": "USD",
                "dedupeCompression": 1.5,
                "discountPercentage": 0,
                "failuresToTolerateAndRaidLevel": "Ftt1Raid1",
                "groupType": "Default",
                "isStretchClusterEnabled": True,
                "memOvercommit": 1,
                "nodeType": "AV36",
                "percentile": "Percentile95",
                "perfDataEndTime": "2023-09-26T13:35:56.5671462Z",
                "perfDataStartTime": "2023-09-25T13:35:56.5671462Z",
                "provisioningState": "Succeeded",
                "reservedInstance": "RI3Year",
                "scalingFactor": 1,
                "sizingCriterion": "AsOnPremises",
                "stage": "InProgress",
                "status": "Created",
                "suitability": "Unknown",
                "suitabilityExplanation": "Unknown",
                "timeRange": "Day",
                "vcpuOversubscription": 4,
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/stable/2023-03-15/examples/AvsAssessmentsOperations_Create_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
