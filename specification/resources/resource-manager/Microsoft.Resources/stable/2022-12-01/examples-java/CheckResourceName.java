
import com.azure.resourcemanager.resources.models.ResourceName;

/** Samples for ResourceName CheckResourceName. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Resources/stable/2022-12-01/examples/CheckResourceName.json
     */
    /**
     * Sample code: CheckValidityForAResourceName.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void checkValidityForAResourceName(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().subscriptionClient().getResourceNames()
            .checkResourceNameWithResponse(new ResourceName(), com.azure.core.util.Context.NONE);
    }
}
