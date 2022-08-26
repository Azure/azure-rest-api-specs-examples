import com.azure.core.util.Context;

/** Samples for Users Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/Users/deleteUser.json
     */
    /**
     * Sample code: deleteUser.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void deleteUser(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.users().delete("testrg123", "testlab", "testuser", Context.NONE);
    }
}
