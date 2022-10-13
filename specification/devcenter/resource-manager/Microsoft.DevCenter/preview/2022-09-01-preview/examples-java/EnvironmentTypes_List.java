import com.azure.core.util.Context;

/** Samples for EnvironmentTypes ListByDevCenter. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-09-01-preview/examples/EnvironmentTypes_List.json
     */
    /**
     * Sample code: EnvironmentTypes_ListByDevCenter.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void environmentTypesListByDevCenter(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.environmentTypes().listByDevCenter("rg1", "Contoso", null, Context.NONE);
    }
}
