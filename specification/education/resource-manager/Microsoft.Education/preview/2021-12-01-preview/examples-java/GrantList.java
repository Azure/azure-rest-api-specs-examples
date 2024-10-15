
/**
 * Samples for Grants ListAll.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/GrantList.json
     */
    /**
     * Sample code: GrantList.
     * 
     * @param manager Entry point to EducationManager.
     */
    public static void grantList(com.azure.resourcemanager.education.EducationManager manager) {
        manager.grants().listAll(false, com.azure.core.util.Context.NONE);
    }
}
