import com.azure.resourcemanager.frontdoor.models.ValidateCustomDomainInput;

/** Samples for FrontDoors ValidateCustomDomain. */
public final class Main {
    /*
     * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/FrontdoorValidateCustomDomain.json
     */
    /**
     * Sample code: FrontDoor_ValidateCustomDomain.
     *
     * @param manager Entry point to FrontDoorManager.
     */
    public static void frontDoorValidateCustomDomain(com.azure.resourcemanager.frontdoor.FrontDoorManager manager) {
        manager
            .frontDoors()
            .validateCustomDomainWithResponse(
                "rg1",
                "frontDoor1",
                new ValidateCustomDomainInput().withHostname("www.someDomain.com"),
                com.azure.core.util.Context.NONE);
    }
}
