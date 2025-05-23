from azure.identity import DefaultAzureCredential

from azure.mgmt.migrationassessment import MigrationAssessmentMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-migrationassessment
# USAGE
    python assessments_operations_create_maximum_set_gen.py

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

    response = client.assessments_operations.begin_create(
        resource_group_name="ayagrawrg",
        project_name="app18700project",
        group_name="kuchatur-test",
        assessment_name="asm1",
        resource={
            "properties": {
                "assessmentType": "Unknown",
                "azureDiskTypes": ["Premium", "StandardSSD"],
                "azureHybridUseBenefit": "Unknown",
                "azureLocation": "njxbwdtsxzhichsnk",
                "azureOfferCode": "Unknown",
                "azurePricingTier": "Standard",
                "azureStorageRedundancy": "Unknown",
                "azureVmFamilies": ["D_series", "Lsv2_series", "M_series", "Mdsv2_series", "Msv2_series", "Mv2_series"],
                "currency": "Unknown",
                "discountPercentage": 6,
                "eaSubscriptionId": "kwsu",
                "groupType": "Default",
                "percentile": "Percentile50",
                "perfDataEndTime": "2023-09-26T09:36:48.491Z",
                "perfDataStartTime": "2023-09-26T09:36:48.491Z",
                "provisioningState": "Succeeded",
                "reservedInstance": "None",
                "scalingFactor": 24,
                "sizingCriterion": "PerformanceBased",
                "stage": "InProgress",
                "status": "Created",
                "timeRange": "Day",
                "vmUptime": {"daysPerMonth": 13, "hoursPerDay": 26},
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/stable/2023-03-15/examples/AssessmentsOperations_Create_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
