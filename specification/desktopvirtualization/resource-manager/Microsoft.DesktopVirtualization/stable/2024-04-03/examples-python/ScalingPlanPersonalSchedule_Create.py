from azure.identity import DefaultAzureCredential

from azure.mgmt.desktopvirtualization import DesktopVirtualizationMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-desktopvirtualization
# USAGE
    python scaling_plan_personal_schedule_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DesktopVirtualizationMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="daefabc0-95b4-48b3-b645-8a753a63c4fa",
    )

    response = client.scaling_plan_personal_schedules.create(
        resource_group_name="resourceGroup1",
        scaling_plan_name="scalingPlan1",
        scaling_plan_schedule_name="scalingPlanScheduleWeekdays1",
        scaling_plan_schedule={
            "properties": {
                "daysOfWeek": ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday"],
                "offPeakActionOnDisconnect": "None",
                "offPeakActionOnLogoff": "Deallocate",
                "offPeakMinutesToWaitOnDisconnect": 10,
                "offPeakMinutesToWaitOnLogoff": 10,
                "offPeakStartTime": {"hour": 20, "minute": 0},
                "offPeakStartVMOnConnect": "Enable",
                "peakActionOnDisconnect": "None",
                "peakActionOnLogoff": "Deallocate",
                "peakMinutesToWaitOnDisconnect": 10,
                "peakMinutesToWaitOnLogoff": 10,
                "peakStartTime": {"hour": 8, "minute": 0},
                "peakStartVMOnConnect": "Enable",
                "rampDownActionOnDisconnect": "None",
                "rampDownActionOnLogoff": "Deallocate",
                "rampDownMinutesToWaitOnDisconnect": 10,
                "rampDownMinutesToWaitOnLogoff": 10,
                "rampDownStartTime": {"hour": 18, "minute": 0},
                "rampDownStartVMOnConnect": "Enable",
                "rampUpActionOnDisconnect": "None",
                "rampUpActionOnLogoff": "None",
                "rampUpAutoStartHosts": "All",
                "rampUpMinutesToWaitOnDisconnect": 10,
                "rampUpMinutesToWaitOnLogoff": 10,
                "rampUpStartTime": {"hour": 6, "minute": 0},
                "rampUpStartVMOnConnect": "Enable",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/ScalingPlanPersonalSchedule_Create.json
if __name__ == "__main__":
    main()
