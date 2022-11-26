import com.azure.core.util.Context;

/** Samples for EnvironmentTypes Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/EnvironmentTypes_Delete.json
     */
    /**
     * Sample code: EnvironmentTypes_Delete.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void environmentTypesDelete(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.environmentTypes().deleteWithResponse("rg1", "Contoso", "DevTest", Context.NONE);
    }
}
