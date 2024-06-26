import com.azure.core.util.Context;

/** Samples for Galleries ListByDevCenter. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/Galleries_List.json
     */
    /**
     * Sample code: Galleries_ListByDevCenter.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void galleriesListByDevCenter(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.galleries().listByDevCenter("rg1", "Contoso", null, Context.NONE);
    }
}
