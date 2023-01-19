/** Samples for Applications DeleteById. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/deleteApplicationById.json
     */
    /**
     * Sample code: Delete application by id.
     *
     * @param manager Entry point to ApplicationManager.
     */
    public static void deleteApplicationById(com.azure.resourcemanager.managedapplications.ApplicationManager manager) {
        manager.applications().deleteById("myApplicationId", com.azure.core.util.Context.NONE);
    }
}
