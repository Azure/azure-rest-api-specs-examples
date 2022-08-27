import com.azure.core.util.Context;

/** Samples for Users ListByLab. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/Users/listUser.json
     */
    /**
     * Sample code: listUser.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void listUser(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.users().listByLab("testrg123", "testlab", null, Context.NONE);
    }
}
