
import com.azure.resourcemanager.resources.fluent.models.DataBoundaryDefinitionInner;
import com.azure.resourcemanager.resources.models.DataBoundary;
import com.azure.resourcemanager.resources.models.DataBoundaryProperties;
import com.azure.resourcemanager.resources.models.DefaultName;

/**
 * Samples for DataBoundaries Put.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Resources/stable/2024-08-01/examples/PutDataBoundary.json
     */
    /**
     * Sample code: Opt-in to data boundary.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void optInToDataBoundary(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().dataBoundaryClient().getDataBoundaries().putWithResponse(DefaultName.DEFAULT,
            new DataBoundaryDefinitionInner().withProperties(
                new DataBoundaryProperties().withDataBoundary(DataBoundary.EU)),
            com.azure.core.util.Context.NONE);
    }
}
