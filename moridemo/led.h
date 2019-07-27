#pragma once

#define LED_WIDTH 30
#define LED_HEIGHT 100
#define LED_DEPTH 30

#ifdef __cplusplus
extern "C"
{
#endif /* __cplusplus */

    /*!
 Specify the URL of 3D LED CUBE.
 If you specify a URL, the drawing data is transferred each time you call the
 Show function.
 @param[in] url 3D LED CUBE IP ADDRESS
 @param[in] port 3D LED CUBE PORT
 @return 0: success
*/
    int InitSdk(char const *url, char const *port);
    void DeinitSdk();
    /*!
 Specify the color of each LED.
 Just by calling this function, drawing data is not transferred to 3D LED CUBE.
 @param[in] x 0 <= x < LED_WIDTH
 @param[in] y 0 <= y < LED_HEIGHT
 @param[in] z 0 <= z < LED_DEPTH
 @param[in] rgb 0x000000 <= rgb <= 0xFFFFFF
*/
    void SetLed(int x, int y, int z, int rgb);

    /*!
 Clear all LED.
 Just by calling this function, drawing data is not transferred to 3D LED CUBE.
*/
    void Clear();

    /*!
 Send packet to 3D LED
 @return 0: success
*/
    int Show();

    /*!
 Wait for the specified time.
 This function stores the time waiting for the previous time,
 taking into consideration the processing delay.
*/
    void Wait(int ms);

#ifdef __cplusplus
}
#endif /* __cplusplus */
