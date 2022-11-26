import com.azure.core.util.Context;

/** Samples for Images ListByDevCenter. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/Images_ListByDevCenter.json
     */
    /**
     * Sample code: Images_ListByDevCenter.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void imagesListByDevCenter(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.images().listByDevCenter("rg1", "Contoso", null, Context.NONE);
    }
}
