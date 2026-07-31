
import com.azure.resourcemanager.appservice.models.VnetParameters;

/**
 * Samples for ResourceProviders VerifyHostingEnvironmentVnet.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/VerifyHostingEnvironmentVnet.json
     */
    /**
     * Sample code: VerifyHostingEnvironmentVnet.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void verifyHostingEnvironmentVnet(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getResourceProviders().verifyHostingEnvironmentVnetWithResponse(new VnetParameters()
            .withVnetResourceGroup("vNet123rg").withVnetName("vNet123").withVnetSubnetName("vNet123SubNet"),
            com.azure.core.util.Context.NONE);
    }
}
