import com.azure.core.util.Context;

/** Samples for Assessments Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/Assessments/DeleteAssessment_example.json
     */
    /**
     * Sample code: Delete a security recommendation task on a resource.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void deleteASecurityRecommendationTaskOnAResource(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .assessments()
            .deleteByResourceGroupWithResponse(
                "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss2",
                "8bb8be0a-6010-4789-812f-e4d661c4ed0e",
                Context.NONE);
    }
}
