
/**
 * Samples for CustomEntityStoreAssignments GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/
     * CustomEntityStoreAssignments/customEntityStoreAssignmentGet_example.json
     */
    /**
     * Sample code: Get a custom entity store assignment.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getACustomEntityStoreAssignment(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.customEntityStoreAssignments().getByResourceGroupWithResponse("TestResourceGroup",
            "33e7cc6e-a139-4723-a0e5-76993aee0771", com.azure.core.util.Context.NONE);
    }
}
