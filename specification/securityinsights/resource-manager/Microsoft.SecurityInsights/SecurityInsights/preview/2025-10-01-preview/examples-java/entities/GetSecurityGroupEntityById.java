
/**
 * Samples for Entities Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/entities/GetSecurityGroupEntityById.json
     */
    /**
     * Sample code: Get a security group entity.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getASecurityGroupEntity(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.entities().getWithResponse("myRg", "myWorkspace", "e1d3d618-e11f-478b-98e3-bb381539a8e1",
            com.azure.core.util.Context.NONE);
    }
}
