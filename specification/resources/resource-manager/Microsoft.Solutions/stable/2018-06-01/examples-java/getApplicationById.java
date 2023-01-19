/** Samples for Applications GetById. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/getApplicationById.json
     */
    /**
     * Sample code: Get application by id.
     *
     * @param manager Entry point to ApplicationManager.
     */
    public static void getApplicationById(com.azure.resourcemanager.managedapplications.ApplicationManager manager) {
        manager.applications().getByIdWithResponse("myApplicationId", com.azure.core.util.Context.NONE);
    }
}
