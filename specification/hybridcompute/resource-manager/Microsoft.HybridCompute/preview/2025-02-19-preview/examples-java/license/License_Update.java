
import com.azure.resourcemanager.hybridcompute.models.License;
import com.azure.resourcemanager.hybridcompute.models.LicenseCoreType;
import com.azure.resourcemanager.hybridcompute.models.LicenseEdition;
import com.azure.resourcemanager.hybridcompute.models.LicenseState;
import com.azure.resourcemanager.hybridcompute.models.LicenseTarget;
import com.azure.resourcemanager.hybridcompute.models.LicenseType;

/**
 * Samples for Licenses Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/license/
     * License_Update.json
     */
    /**
     * Sample code: Update a License.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void updateALicense(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        License resource = manager.licenses()
            .getByResourceGroupWithResponse("myResourceGroup", "{licenseName}", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withLicenseType(LicenseType.ESU).withState(LicenseState.ACTIVATED)
            .withTarget(LicenseTarget.WINDOWS_SERVER_2012).withEdition(LicenseEdition.DATACENTER)
            .withType(LicenseCoreType.P_CORE).withProcessors(6).apply();
    }
}
