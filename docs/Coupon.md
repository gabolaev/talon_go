# Coupon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**CampaignId** | Pointer to **int32** | The ID of the campaign that owns this entity. | 
**Value** | Pointer to **string** | The coupon code. | 
**UsageLimit** | Pointer to **int32** | The number of times the coupon code can be redeemed. &#x60;0&#x60; means unlimited redemptions but any campaign usage limits will still apply.  | 
**DiscountLimit** | Pointer to **float32** | The total discount value that the code can give. Typically used to represent a gift card value.  | [optional] 
**ReservationLimit** | Pointer to **int32** | The number of reservations that can be made with this coupon code.  | [optional] 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the coupon becomes valid. | [optional] 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | Expiration date of the coupon. Coupon never expires if this is omitted, zero, or negative. | [optional] 
**Limits** | Pointer to [**[]LimitConfig**](LimitConfig.md) | Limits configuration for a coupon. These limits will override the limits set from the campaign.  **Note:** Only usable when creating a single coupon which is not tied to a specific recipient. Only per-profile limits are allowed to be configured.  | [optional] 
**UsageCounter** | Pointer to **int32** | The number of times the coupon has been successfully redeemed. | 
**DiscountCounter** | Pointer to **float32** | The amount of discounts given on rules redeeming this coupon. Only usable if a coupon discount budget was set for this coupon. | [optional] 
**DiscountRemainder** | Pointer to **float32** | The remaining discount this coupon can give. | [optional] 
**ReservationCounter** | Pointer to **float32** | The number of times this coupon has been reserved. | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Custom attributes associated with this coupon. | [optional] 
**ReferralId** | Pointer to **int32** | The integration ID of the referring customer (if any) for whom this coupon was created as an effect. | [optional] 
**RecipientIntegrationId** | Pointer to **string** | The Integration ID of the customer that is allowed to redeem this coupon. | [optional] 
**ImportId** | Pointer to **int32** | The ID of the Import which created this coupon. | [optional] 
**Reservation** | Pointer to **bool** | Defines the type of reservation: - &#x60;true&#x60;: The reservation is a soft reservation. Any customer can use the coupon. This is done via the [Create coupon reservation](https://docs.talon.one/integration-api#operation/createCouponReservation) endpoint. - &#x60;false&#x60;: The reservation is a hard reservation. Only the associated customer (&#x60;recipientIntegrationId&#x60;) can use the coupon. This is done via the Campaign Manager when you create a coupon for a given &#x60;recipientIntegrationId&#x60;, the [Create coupons](https://docs.talon.one/management-api#operation/createCoupons) endpoint or [Create coupons for multiple recipients](https://docs.talon.one/management-api#operation/createCouponsForMultipleRecipients) endpoint.  | [optional] [default to true]
**BatchId** | Pointer to **string** | The id of the batch the coupon belongs to. | [optional] 
**IsReservationMandatory** | Pointer to **bool** | Whether the reservation effect actually created a new reservation. | [optional] [default to true]

## Methods

### GetId

`func (o *Coupon) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Coupon) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Coupon) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Coupon) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *Coupon) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Coupon) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Coupon) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Coupon) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetCampaignId

`func (o *Coupon) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *Coupon) GetCampaignIdOk() (int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignId

`func (o *Coupon) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignId

`func (o *Coupon) SetCampaignId(v int32)`

SetCampaignId gets a reference to the given int32 and assigns it to the CampaignId field.

### GetValue

`func (o *Coupon) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Coupon) GetValueOk() (string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValue

`func (o *Coupon) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValue

`func (o *Coupon) SetValue(v string)`

SetValue gets a reference to the given string and assigns it to the Value field.

### GetUsageLimit

`func (o *Coupon) GetUsageLimit() int32`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *Coupon) GetUsageLimitOk() (int32, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsageLimit

`func (o *Coupon) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### SetUsageLimit

`func (o *Coupon) SetUsageLimit(v int32)`

SetUsageLimit gets a reference to the given int32 and assigns it to the UsageLimit field.

### GetDiscountLimit

`func (o *Coupon) GetDiscountLimit() float32`

GetDiscountLimit returns the DiscountLimit field if non-nil, zero value otherwise.

### GetDiscountLimitOk

`func (o *Coupon) GetDiscountLimitOk() (float32, bool)`

GetDiscountLimitOk returns a tuple with the DiscountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiscountLimit

`func (o *Coupon) HasDiscountLimit() bool`

HasDiscountLimit returns a boolean if a field has been set.

### SetDiscountLimit

`func (o *Coupon) SetDiscountLimit(v float32)`

SetDiscountLimit gets a reference to the given float32 and assigns it to the DiscountLimit field.

### GetReservationLimit

`func (o *Coupon) GetReservationLimit() int32`

GetReservationLimit returns the ReservationLimit field if non-nil, zero value otherwise.

### GetReservationLimitOk

`func (o *Coupon) GetReservationLimitOk() (int32, bool)`

GetReservationLimitOk returns a tuple with the ReservationLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReservationLimit

`func (o *Coupon) HasReservationLimit() bool`

HasReservationLimit returns a boolean if a field has been set.

### SetReservationLimit

`func (o *Coupon) SetReservationLimit(v int32)`

SetReservationLimit gets a reference to the given int32 and assigns it to the ReservationLimit field.

### GetStartDate

`func (o *Coupon) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Coupon) GetStartDateOk() (time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartDate

`func (o *Coupon) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDate

`func (o *Coupon) SetStartDate(v time.Time)`

SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.

### GetExpiryDate

`func (o *Coupon) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *Coupon) GetExpiryDateOk() (time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiryDate

`func (o *Coupon) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDate

`func (o *Coupon) SetExpiryDate(v time.Time)`

SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.

### GetLimits

`func (o *Coupon) GetLimits() []LimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *Coupon) GetLimitsOk() ([]LimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLimits

`func (o *Coupon) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### SetLimits

`func (o *Coupon) SetLimits(v []LimitConfig)`

SetLimits gets a reference to the given []LimitConfig and assigns it to the Limits field.

### GetUsageCounter

`func (o *Coupon) GetUsageCounter() int32`

GetUsageCounter returns the UsageCounter field if non-nil, zero value otherwise.

### GetUsageCounterOk

`func (o *Coupon) GetUsageCounterOk() (int32, bool)`

GetUsageCounterOk returns a tuple with the UsageCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsageCounter

`func (o *Coupon) HasUsageCounter() bool`

HasUsageCounter returns a boolean if a field has been set.

### SetUsageCounter

`func (o *Coupon) SetUsageCounter(v int32)`

SetUsageCounter gets a reference to the given int32 and assigns it to the UsageCounter field.

### GetDiscountCounter

`func (o *Coupon) GetDiscountCounter() float32`

GetDiscountCounter returns the DiscountCounter field if non-nil, zero value otherwise.

### GetDiscountCounterOk

`func (o *Coupon) GetDiscountCounterOk() (float32, bool)`

GetDiscountCounterOk returns a tuple with the DiscountCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiscountCounter

`func (o *Coupon) HasDiscountCounter() bool`

HasDiscountCounter returns a boolean if a field has been set.

### SetDiscountCounter

`func (o *Coupon) SetDiscountCounter(v float32)`

SetDiscountCounter gets a reference to the given float32 and assigns it to the DiscountCounter field.

### GetDiscountRemainder

`func (o *Coupon) GetDiscountRemainder() float32`

GetDiscountRemainder returns the DiscountRemainder field if non-nil, zero value otherwise.

### GetDiscountRemainderOk

`func (o *Coupon) GetDiscountRemainderOk() (float32, bool)`

GetDiscountRemainderOk returns a tuple with the DiscountRemainder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiscountRemainder

`func (o *Coupon) HasDiscountRemainder() bool`

HasDiscountRemainder returns a boolean if a field has been set.

### SetDiscountRemainder

`func (o *Coupon) SetDiscountRemainder(v float32)`

SetDiscountRemainder gets a reference to the given float32 and assigns it to the DiscountRemainder field.

### GetReservationCounter

`func (o *Coupon) GetReservationCounter() float32`

GetReservationCounter returns the ReservationCounter field if non-nil, zero value otherwise.

### GetReservationCounterOk

`func (o *Coupon) GetReservationCounterOk() (float32, bool)`

GetReservationCounterOk returns a tuple with the ReservationCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReservationCounter

`func (o *Coupon) HasReservationCounter() bool`

HasReservationCounter returns a boolean if a field has been set.

### SetReservationCounter

`func (o *Coupon) SetReservationCounter(v float32)`

SetReservationCounter gets a reference to the given float32 and assigns it to the ReservationCounter field.

### GetAttributes

`func (o *Coupon) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Coupon) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *Coupon) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *Coupon) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetReferralId

`func (o *Coupon) GetReferralId() int32`

GetReferralId returns the ReferralId field if non-nil, zero value otherwise.

### GetReferralIdOk

`func (o *Coupon) GetReferralIdOk() (int32, bool)`

GetReferralIdOk returns a tuple with the ReferralId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralId

`func (o *Coupon) HasReferralId() bool`

HasReferralId returns a boolean if a field has been set.

### SetReferralId

`func (o *Coupon) SetReferralId(v int32)`

SetReferralId gets a reference to the given int32 and assigns it to the ReferralId field.

### GetRecipientIntegrationId

`func (o *Coupon) GetRecipientIntegrationId() string`

GetRecipientIntegrationId returns the RecipientIntegrationId field if non-nil, zero value otherwise.

### GetRecipientIntegrationIdOk

`func (o *Coupon) GetRecipientIntegrationIdOk() (string, bool)`

GetRecipientIntegrationIdOk returns a tuple with the RecipientIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecipientIntegrationId

`func (o *Coupon) HasRecipientIntegrationId() bool`

HasRecipientIntegrationId returns a boolean if a field has been set.

### SetRecipientIntegrationId

`func (o *Coupon) SetRecipientIntegrationId(v string)`

SetRecipientIntegrationId gets a reference to the given string and assigns it to the RecipientIntegrationId field.

### GetImportId

`func (o *Coupon) GetImportId() int32`

GetImportId returns the ImportId field if non-nil, zero value otherwise.

### GetImportIdOk

`func (o *Coupon) GetImportIdOk() (int32, bool)`

GetImportIdOk returns a tuple with the ImportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImportId

`func (o *Coupon) HasImportId() bool`

HasImportId returns a boolean if a field has been set.

### SetImportId

`func (o *Coupon) SetImportId(v int32)`

SetImportId gets a reference to the given int32 and assigns it to the ImportId field.

### GetReservation

`func (o *Coupon) GetReservation() bool`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *Coupon) GetReservationOk() (bool, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReservation

`func (o *Coupon) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### SetReservation

`func (o *Coupon) SetReservation(v bool)`

SetReservation gets a reference to the given bool and assigns it to the Reservation field.

### GetBatchId

`func (o *Coupon) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *Coupon) GetBatchIdOk() (string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBatchId

`func (o *Coupon) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### SetBatchId

`func (o *Coupon) SetBatchId(v string)`

SetBatchId gets a reference to the given string and assigns it to the BatchId field.

### GetIsReservationMandatory

`func (o *Coupon) GetIsReservationMandatory() bool`

GetIsReservationMandatory returns the IsReservationMandatory field if non-nil, zero value otherwise.

### GetIsReservationMandatoryOk

`func (o *Coupon) GetIsReservationMandatoryOk() (bool, bool)`

GetIsReservationMandatoryOk returns a tuple with the IsReservationMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsReservationMandatory

`func (o *Coupon) HasIsReservationMandatory() bool`

HasIsReservationMandatory returns a boolean if a field has been set.

### SetIsReservationMandatory

`func (o *Coupon) SetIsReservationMandatory(v bool)`

SetIsReservationMandatory gets a reference to the given bool and assigns it to the IsReservationMandatory field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


