
import com.azure.resourcemanager.hybridnetwork.models.NetworkFunctionValueWithoutSecrets;
import com.azure.resourcemanager.hybridnetwork.models.NfviType;
import com.azure.resourcemanager.hybridnetwork.models.OpenDeploymentResourceReference;
import java.util.Arrays;

/**
 * Samples for NetworkFunctions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * NetworkFunctionCreate.json
     */
    /**
     * Sample code: Create network function resource.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void
        createNetworkFunctionResource(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.networkFunctions().define("testNf").withRegion("eastus").withExistingResourceGroup("rg")
            .withProperties(new NetworkFunctionValueWithoutSecrets()
                .withNetworkFunctionDefinitionVersionResourceReference(new OpenDeploymentResourceReference().withId(
                    "/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/publishers/testVendor/networkFunctionDefinitionGroups/testnetworkFunctionDefinitionGroupName/networkFunctionDefinitionVersions/1.0.1"))
                .withNfviType(NfviType.AZURE_ARC_KUBERNETES)
                .withNfviId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testResourceGroup/providers/Microsoft.ExtendedLocation/customLocations/testCustomLocation")
                .withAllowSoftwareUpdate(false)
                .withRoleOverrideValues(Arrays.asList(
                    "{\"name\":\"testRoleOne\",\"deployParametersMappingRuleProfile\":{\"helmMappingRuleProfile\":{\"helmPackageVersion\":\"2.1.3\",\"values\":\"{\\\"roleOneParam\\\":\\\"roleOneOverrideValue\\\"}\"}}}",
                    "{\"name\":\"testRoleTwo\",\"deployParametersMappingRuleProfile\":{\"helmMappingRuleProfile\":{\"releaseName\":\"overrideReleaseName\",\"releaseNamespace\":\"overrideNamespace\",\"values\":\"{\\\"roleTwoParam\\\":\\\"roleTwoOverrideValue\\\"}\"}}}"))
                .withDeploymentValues("{\"releaseName\":\"testReleaseName\",\"namespace\":\"testNamespace\"}"))
            .create();
    }
}
