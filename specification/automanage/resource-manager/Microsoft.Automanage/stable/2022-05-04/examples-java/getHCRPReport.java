import com.azure.core.util.Context;

/** Samples for HcrpReports Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/getHCRPReport.json
     */
    /**
     * Sample code: Get a report for a HCRP configuration profile assignment.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void getAReportForAHCRPConfigurationProfileAssignment(
        com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager
            .hcrpReports()
            .getWithResponse(
                "myResourceGroupName",
                "myMachineName",
                "default",
                "b4e9ee6b-1717-4ff0-a8d2-e6d72c33d5f4",
                Context.NONE);
    }
}
