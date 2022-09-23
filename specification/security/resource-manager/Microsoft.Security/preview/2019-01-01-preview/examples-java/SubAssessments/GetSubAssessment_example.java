import com.azure.core.util.Context;

/** Samples for SubAssessments Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/SubAssessments/GetSubAssessment_example.json
     */
    /**
     * Sample code: Get security recommendation task from security data location.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void getSecurityRecommendationTaskFromSecurityDataLocation(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .subAssessments()
            .getWithResponse(
                "subscriptions/212f9889-769e-45ae-ab43-6da33674bd26/resourceGroups/DEMORG/providers/Microsoft.Compute/virtualMachines/vm2",
                "1195afff-c881-495e-9bc5-1486211ae03f",
                "95f7da9c-a2a4-1322-0758-fcd24ef09b85",
                Context.NONE);
    }
}
